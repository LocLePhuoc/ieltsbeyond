import type { Assessment, AssessmentCriterion } from "../../lib/api";

const CRITERIA_LABELS: Record<keyof Assessment["criteria"], string> = {
  task_achievement: "Task Achievement",
  coherence_and_cohesion: "Coherence & Cohesion",
  lexical_resource: "Lexical Resource",
  grammatical_range_accuracy: "Grammatical Range & Accuracy",
};

interface AssessmentPanelProps {
  assessing: boolean;
  assessError: string | null;
  assessment: Assessment | null;
}

function CriterionCard({ label, criterion }: { label: string; criterion: AssessmentCriterion }) {
  return (
    <div className="rounded-xl border border-white/60 bg-white/60 p-3.5">
      <div className="flex items-center justify-between mb-1.5">
        <span className="text-[13px] font-semibold text-charcoal">{label}</span>
        <span className="text-[13px] font-bold text-sage">{criterion.band.toFixed(1)}</span>
      </div>
      {criterion.comment && <p className="text-[12.5px] text-charcoal-light/80 leading-relaxed mb-2">{criterion.comment}</p>}
      {criterion.strengths?.length > 0 && (
        <ul className="mb-1.5 space-y-0.5">
          {criterion.strengths.map((s, i) => (
            <li key={i} className="text-[12px] text-charcoal-light/80 leading-snug flex gap-1.5">
              <span className="text-sage">+</span>
              <span>{s}</span>
            </li>
          ))}
        </ul>
      )}
      {criterion.weaknesses?.length > 0 && (
        <ul className="space-y-0.5">
          {criterion.weaknesses.map((w, i) => (
            <li key={i} className="text-[12px] text-charcoal-light/80 leading-snug flex gap-1.5">
              <span className="text-red-400">–</span>
              <span>{w}</span>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default function AssessmentPanel({ assessing, assessError, assessment }: AssessmentPanelProps) {
  if (assessing) {
    return (
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-6 flex flex-col items-center justify-center text-center gap-3 min-h-[240px]">
        <div className="h-6 w-6 rounded-full border-2 border-sage border-t-transparent animate-spin" />
        <p className="text-sm text-charcoal-light/70">Assessing your writing...</p>
      </section>
    );
  }

  if (assessError) {
    return (
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-6">
        <p className="text-sm text-red-500">{assessError}</p>
      </section>
    );
  }

  if (!assessment) {
    return (
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-6 min-h-[240px] flex items-center justify-center">
        <p className="text-sm text-charcoal-light/50 text-center">
          Your band scores and feedback will appear here after you check your answer.
        </p>
      </section>
    );
  }

  const bands = Object.values(assessment.criteria).map((c) => c.band);
  const overallBand = bands.length > 0 ? bands.reduce((a, b) => a + b, 0) / bands.length : 0;

  return (
    <div className="flex flex-col gap-4">
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-5">
        <div className="flex items-center justify-between mb-3">
          <span className="text-[11px] font-semibold tracking-wide uppercase text-charcoal-light/60">Overall Band</span>
          <span className="text-2xl font-bold text-sage">{overallBand.toFixed(1)}</span>
        </div>
        {assessment.overview_feedback && (
          <p className="text-[13px] text-charcoal-light/80 leading-relaxed">{assessment.overview_feedback}</p>
        )}
      </section>

      <div className="flex flex-col gap-2.5">
        {(Object.keys(CRITERIA_LABELS) as Array<keyof Assessment["criteria"]>).map((key) => (
          <CriterionCard key={key} label={CRITERIA_LABELS[key]} criterion={assessment.criteria[key]} />
        ))}
      </div>

      {(assessment.data_accuracy_check?.correct_points?.length > 0 ||
        assessment.data_accuracy_check?.incorrect_points?.length > 0) && (
        <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-4">
          <h3 className="text-[13px] font-semibold text-charcoal mb-2">Data Accuracy</h3>
          {assessment.data_accuracy_check.correct_points?.map((p, i) => (
            <p key={`c-${i}`} className="text-[12px] text-charcoal-light/80 leading-snug flex gap-1.5 mb-1">
              <span className="text-sage">✓</span>
              <span>{p}</span>
            </p>
          ))}
          {assessment.data_accuracy_check.incorrect_points?.map((p, i) => (
            <p key={`i-${i}`} className="text-[12px] text-charcoal-light/80 leading-snug flex gap-1.5 mb-1">
              <span className="text-red-400">✗</span>
              <span>{p}</span>
            </p>
          ))}
        </section>
      )}

      {assessment.error_corrections?.length > 0 && (
        <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-4">
          <h3 className="text-[13px] font-semibold text-charcoal mb-2">Error Corrections</h3>
          <div className="flex flex-col gap-2.5">
            {assessment.error_corrections.map((e, i) => (
              <div key={i} className="text-[12px] leading-snug border-l-2 border-sage/40 pl-2.5">
                <p className="text-charcoal-light/60 line-through">{e.original}</p>
                <p className="text-charcoal font-medium">{e.correction}</p>
                {e.explanation && <p className="text-charcoal-light/70 mt-0.5">{e.explanation}</p>}
              </div>
            ))}
          </div>
        </section>
      )}

      {assessment.top_priorities_to_improve?.length > 0 && (
        <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-4">
          <h3 className="text-[13px] font-semibold text-charcoal mb-2">Top Priorities to Improve</h3>
          <ol className="list-decimal list-inside space-y-1">
            {assessment.top_priorities_to_improve.map((p, i) => (
              <li key={i} className="text-[12px] text-charcoal-light/80 leading-snug">
                {p}
              </li>
            ))}
          </ol>
        </section>
      )}
    </div>
  );
}

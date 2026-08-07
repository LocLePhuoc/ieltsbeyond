import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { assessWritingTask1, submitWritingTask1, type Assessment } from "../../lib/api";
import AssessmentPanel from "./AssessmentPanel";

interface PracticeTask {
  id?: string;
  type: string;
  category: string;
  question: string;
  imageKey?: string;
}

const defaultTask: PracticeTask = {
  type: "Task 2",
  category: "Opinion",
  question:
    "Some people think that the best way to increase road safety is to raise the minimum legal age for driving cars or riding motorbikes. To what extent do you agree or disagree?",
};

export default function PracticePage() {
  const location = useLocation();
  const task = (location.state as PracticeTask | null) ?? defaultTask;
  const [answer, setAnswer] = useState("");
  const [imageFailed, setImageFailed] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);
  const [assessing, setAssessing] = useState(false);
  const [assessError, setAssessError] = useState<string | null>(null);
  const [assessment, setAssessment] = useState<Assessment | null>(null);

  const isTask1 = task.type === "Task 1";

  useEffect(() => setImageFailed(false), [task.imageKey]);

  async function handleCheckAnswer() {
    if (!isTask1 || !task.id) return;

    setSubmitting(true);
    setSubmitError(null);
    setAssessError(null);
    setAssessment(null);
    try {
      const submission = await submitWritingTask1(task.id, answer);
      setSubmitting(false);

      setAssessing(true);
      try {
        const result = await assessWritingTask1(task.id, submission.id);
        setAssessment(result);
      } catch {
        setAssessError("Failed to assess your answer. Please try again.");
      } finally {
        setAssessing(false);
      }
    } catch {
      setSubmitError("Failed to submit your answer. Please try again.");
      setSubmitting(false);
    }
  }

  const editor = (
    <div className="flex flex-col gap-5">
      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-6 md:p-8">
        <div className="flex items-center gap-2 mb-3">
          <span className="text-[11px] font-semibold tracking-wide uppercase text-sage bg-sage-50 px-2.5 py-1 rounded-full">
            {task.type}
          </span>
          <span className="text-[11px] font-medium tracking-wide uppercase text-charcoal-light/60">
            {task.category}
          </span>
        </div>
        <p className="font-serif text-lg md:text-xl text-charcoal leading-relaxed">{task.question}</p>
      </section>

      {task.type === "Task 1" && task.imageKey ? (
        <section className="rounded-2xl border border-white/60 bg-white/40 overflow-hidden">
          {imageFailed ? (
            <div className="flex flex-col items-center justify-center py-14 text-charcoal-light/50 text-sm">
              Image unavailable
            </div>
          ) : (
            <img
              src={task.imageKey}
              alt={`${task.category} chart for this Writing Task 1 question`}
              className="w-full max-h-[420px] object-contain bg-white"
              onError={() => setImageFailed(true)}
            />
          )}
        </section>
      ) : (
        <section className="min-h-[100px] rounded-2xl border border-dashed border-gray-300 bg-white/40" />
      )}

      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-2">
        <textarea
          value={answer}
          onChange={(e) => setAnswer(e.target.value)}
          placeholder="Start writing your response here..."
          className="w-full min-h-[360px] resize-y rounded-2xl bg-transparent p-4 text-[15px] text-charcoal leading-relaxed focus:outline-none"
        />
      </section>

      <div className="flex items-center justify-end gap-3">
        {submitError && <span className="text-sm text-red-500">{submitError}</span>}
        <button
          type="button"
          onClick={handleCheckAnswer}
          disabled={submitting || assessing}
          className="px-6 py-3 rounded-xl bg-sage text-white text-sm font-semibold shadow-soft-lg hover:bg-sage-dark transition-colors disabled:opacity-60 disabled:cursor-wait"
        >
          {submitting ? "Submitting..." : assessing ? "Assessing..." : "Check Answer"}
        </button>
      </div>
    </div>
  );

  if (!isTask1) {
    return editor;
  }

  return (
    <div className="flex flex-col lg:flex-row gap-5">
      <div className="flex-1 min-w-0">{editor}</div>
      <aside className="w-full lg:w-[360px] lg:shrink-0">
        <div className="lg:sticky lg:top-6">
          <AssessmentPanel assessing={assessing} assessError={assessError} assessment={assessment} />
        </div>
      </aside>
    </div>
  );
}

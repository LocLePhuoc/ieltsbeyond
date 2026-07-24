import { useState } from "react";

const task = {
  type: "Task 2",
  category: "Opinion",
  question:
    "Some people think that the best way to increase road safety is to raise the minimum legal age for driving cars or riding motorbikes. To what extent do you agree or disagree?",
};

export default function PracticePage() {
  const [answer, setAnswer] = useState("");

  function handleCheckAnswer() {}

  return (
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

      <section className="min-h-[100px] rounded-2xl border border-dashed border-gray-300 bg-white/40" />

      <section className="bg-white/80 backdrop-blur-xl rounded-3xl shadow-glass border border-white/60 p-2">
        <textarea
          value={answer}
          onChange={(e) => setAnswer(e.target.value)}
          placeholder="Start writing your response here..."
          className="w-full min-h-[360px] resize-y rounded-2xl bg-transparent p-4 text-[15px] text-charcoal leading-relaxed focus:outline-none"
        />
      </section>

      <div className="flex justify-end">
        <button
          type="button"
          onClick={handleCheckAnswer}
          className="px-6 py-3 rounded-xl bg-sage text-white text-sm font-semibold shadow-soft-lg hover:bg-sage-dark transition-colors"
        >
          Check Answer
        </button>
      </div>
    </div>
  );
}

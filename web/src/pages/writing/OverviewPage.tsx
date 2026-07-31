import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  getWritingTask1,
  getWritingTask1List,
  getWritingTask2List,
  type WritingTask1,
  type WritingTask2,
} from "../../lib/api";
import { HorizontalScrollSection } from "../../components/HorizontalScrollSection";
import { WritingTaskCard } from "../../components/WritingTaskCard";

export default function OverviewPage() {
  const navigate = useNavigate();
  const [tasks1, setTasks1] = useState<WritingTask1[] | null>(null);
  const [tasks2, setTasks2] = useState<WritingTask2[] | null>(null);
  const [loadingTask1Id, setLoadingTask1Id] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    getWritingTask1List()
      .then((t1) => !cancelled && setTasks1(t1))
      .catch(() => !cancelled && setTasks1([]));
    getWritingTask2List()
      .then((t2) => !cancelled && setTasks2(t2))
      .catch(() => !cancelled && setTasks2([]));
    return () => {
      cancelled = true;
    };
  }, []);

  async function openTask1(taskId: string) {
    setLoadingTask1Id(taskId);
    try {
      const task = await getWritingTask1(taskId);
      navigate("/writing/practice", {
        state: { type: "Task 1", category: task.type, question: task.question, imageKey: task.image_key },
      });
    } catch {
      setLoadingTask1Id(null);
    }
  }

  function openTask2(task: WritingTask2) {
    navigate("/writing/practice", {
      state: { type: "Task 2", category: task.category, question: task.question },
    });
  }

  return (
    <div className="flex flex-col gap-8">
      <HorizontalScrollSection title="Writing Task 1">
        {tasks1 && tasks1.length > 0
          ? tasks1.map((task) => (
              <WritingTaskCard
                key={task.id}
                badge={task.type}
                question={task.question}
                taskLabel="Task 1"
                onSelect={() => openTask1(task.id)}
                loading={loadingTask1Id === task.id}
              />
            ))
          : tasks1 && (
              <p className="text-charcoal-light/50 text-sm py-6">No Task 1 questions yet.</p>
            )}
      </HorizontalScrollSection>

      <HorizontalScrollSection title="Writing Task 2">
        {tasks2 && tasks2.length > 0
          ? tasks2.map((task) => (
              <WritingTaskCard
                key={task.id}
                badge={task.category}
                question={task.question}
                taskLabel="Task 2"
                onSelect={() => openTask2(task)}
              />
            ))
          : tasks2 && (
              <p className="text-charcoal-light/50 text-sm py-6">No Task 2 questions yet.</p>
            )}
      </HorizontalScrollSection>
    </div>
  );
}

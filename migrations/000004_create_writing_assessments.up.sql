CREATE TABLE IF NOT EXISTS writing_assessments (
  id TEXT PRIMARY KEY,
  task_type TEXT NOT NULL CHECK (task_type IN ('task1', 'task2')),
  task_id TEXT NOT NULL,
  submission_id TEXT NOT NULL REFERENCES writing_submissions (id) ON DELETE CASCADE,
  -- {task_achievement, coherence_and_cohesion, lexical_resource, grammatical_range_accuracy},
  -- each {band, strengths, weaknesses, comment}
  criteria JSONB NOT NULL,
  -- {correct_points, incorrect_points}
  data_accuracy_check JSONB NOT NULL DEFAULT '{"correct_points": [], "incorrect_points": []}'::jsonb,
  error_corrections JSONB NOT NULL DEFAULT '[]'::jsonb,
  overview_feedback TEXT NOT NULL DEFAULT '',
  top_priorities_to_improve JSONB NOT NULL DEFAULT '[]'::jsonb,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS writing_assessments_submission_idx
  ON writing_assessments (submission_id, created_at DESC);

CREATE INDEX IF NOT EXISTS writing_assessments_task_idx
  ON writing_assessments (task_type, task_id);

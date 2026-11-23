-- task_id points to writing_tasks1 or writing_tasks2 depending on task_type,
-- so it cannot be a single foreign key.
CREATE TABLE IF NOT EXISTS writing_submissions (
  id TEXT PRIMARY KEY,
  task_type TEXT NOT NULL CHECK (task_type IN ('task1', 'task2')),
  task_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  paragraphs TEXT[] NOT NULL DEFAULT '{}',
  submit_time TIMESTAMPTZ NOT NULL DEFAULT now(),
  time_taken_seconds INTEGER NOT NULL DEFAULT 0 CHECK (time_taken_seconds >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS writing_submissions_user_submit_time_idx
  ON writing_submissions (user_id, submit_time DESC);

CREATE INDEX IF NOT EXISTS writing_submissions_task_idx
  ON writing_submissions (task_type, task_id);

CREATE TABLE IF NOT EXISTS writing_tasks1 (
  id TEXT PRIMARY KEY,
  question TEXT NOT NULL,
  type TEXT NOT NULL CHECK (type IN ('Line', 'Bar', 'Pie', 'Table', 'Combination', 'Process', 'Map')),
  image_key TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS writing_tasks1_type_idx ON writing_tasks1 (type);

CREATE TABLE IF NOT EXISTS writing_tasks2 (
  id TEXT PRIMARY KEY,
  question TEXT NOT NULL,
  category TEXT NOT NULL CHECK (category IN ('Opinion', 'Discussion', 'Problem And Solution', 'Advantages And Disadvantages', 'Two-Part Question')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS writing_tasks2_category_idx ON writing_tasks2 (category);

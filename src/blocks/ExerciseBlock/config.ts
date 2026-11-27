import type { Block } from 'payload'

export const ExerciseBlock: Block = {
  slug: 'exercise',
  interfaceName: 'ExerciseBlock',
  labels: {
    singular: 'Exercise',
    plural: 'Exercises',
  },
  fields: [
    {
      name: 'question',
      type: 'textarea',
      required: true,
    },
    {
      name: 'choices',
      type: 'array',
      minRows: 2,
      required: true,
      fields: [
        {
          name: 'text',
          type: 'text',
          required: true,
        },
      ],
    },
    {
      name: 'correctAnswer',
      type: 'number',
      required: true,
      min: 0,
      admin: {
        description:
          'Enter the zero-based position of the correct choice. For example, 0 is the first choice.',
      },
    },
    {
      name: 'explanation',
      type: 'textarea',
    },
  ],
}

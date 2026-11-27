'use client'

import { Button } from '@/components/ui/button'
import { useId, useState } from 'react'

type Props = {
  question: string
  choices?: { id?: string | null; text: string }[] | null
  correctAnswer: number
  explanation?: string | null
}

export function ExerciseBlockComponent({
  question,
  choices = [],
  correctAnswer,
  explanation,
}: Props) {
  const [selected, setSelected] = useState<number | null>(null)
  const [reviewed, setReviewed] = useState(false)
  const groupName = useId()

  return (
    <section className="my-6 rounded-lg border border-border bg-card p-4">
      <h3 className="mb-4 font-semibold">{question}</h3>

      <fieldset className="space-y-2" disabled={reviewed}>
        <legend className="sr-only">Choose an answer</legend>

        {choices.map((choice, index) => {
          const isCorrectChoice = index === correctAnswer
          const isSelected = index === selected

          return (
            <label
              key={choice.id ?? index}
              className={`flex cursor-pointer items-start gap-3 rounded-md border p-3 ${
                reviewed && isCorrectChoice
                  ? 'border-green-600 bg-green-500/10'
                  : reviewed && isSelected
                    ? 'border-red-600 bg-red-500/10'
                    : 'border-border'
              }`}
            >
              <input
                type="radio"
                name={groupName}
                checked={isSelected}
                onChange={() => setSelected(index)}
                className="mt-1"
              />
              <span>{choice.text}</span>

              {reviewed && isCorrectChoice && (
                <span className="ml-auto text-sm font-medium text-green-700">Correct answer</span>
              )}
            </label>
          )
        })}
      </fieldset>

      {!reviewed ? (
        <Button className="mt-4" disabled={selected === null} onClick={() => setReviewed(true)}>
          Check answer
        </Button>
      ) : (
        <div className="mt-4 space-y-3" aria-live="polite">
          <p className={selected === correctAnswer ? 'text-green-700' : 'text-red-700'}>
            {selected === correctAnswer
              ? 'Correct!'
              : 'Review the correct answer highlighted above.'}
          </p>

          {explanation && <p className="text-sm text-muted-foreground">{explanation}</p>}

          <Button
            variant="secondary"
            onClick={() => {
              setSelected(null)
              setReviewed(false)
            }}
          >
            Try again
          </Button>
        </div>
      )}
    </section>
  )
}

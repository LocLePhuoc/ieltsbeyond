package model

type WritingTask1Type string

type WritingTask2Category string

const (
	WritingTask1TypeLine        WritingTask1Type = "Line"
	WritingTask1TypeBar         WritingTask1Type = "Bar"
	WritingTask1TypePie         WritingTask1Type = "Pie"
	WritingTask1TypeTable       WritingTask1Type = "Table"
	WritingTask1TypeCombination WritingTask1Type = "Combination"
	WritingTask1TypeProcess     WritingTask1Type = "Process"
	WritingTask1TypeMap         WritingTask1Type = "Map"
)

var AllWritingTask1Types = []WritingTask1Type{
	WritingTask1TypeLine,
	WritingTask1TypeBar,
	WritingTask1TypePie,
	WritingTask1TypeTable,
	WritingTask1TypeCombination,
	WritingTask1TypeProcess,
	WritingTask1TypeMap,
}

func IsValidWritingTask1Type(s string) bool {
	for _, t := range AllWritingTask1Types {
		if string(t) == s {
			return true
		}
	}
	return false
}

const (
	WritingTask2CategoryOpinion                    WritingTask2Category = "Opinion"
	WritingTask2CategoryDiscussion                 WritingTask2Category = "Discussion"
	WritingTask2CategoryProblemAndSolution         WritingTask2Category = "Problem And Solution"
	WritingTask2CategoryAdvantagesAndDisadvantages WritingTask2Category = "Advantages And Disadvantages"
	WritingTask2CategoryTwoPartQuestion            WritingTask2Category = "Two-Part Question"
)

var AllWritingTask2Categories = []WritingTask2Category{
	WritingTask2CategoryOpinion,
	WritingTask2CategoryDiscussion,
	WritingTask2CategoryProblemAndSolution,
	WritingTask2CategoryAdvantagesAndDisadvantages,
	WritingTask2CategoryTwoPartQuestion,
}

func IsValidWritingTask2Category(s string) bool {
	for _, c := range AllWritingTask2Categories {
		if string(c) == s {
			return true
		}
	}
	return false
}

type Paragraph struct {
	Text string
}

type WritingTask1 struct {
	Id                    string
	Question              string
	Type                  WritingTask1Type
	ImagePath             string
	NumRequiredParagraphs int
	Paragraphs            []Paragraph
}

type WritingTask2 struct {
	Id                    string
	Question              string
	Category              WritingTask2Category
	NumRequiredParagraphs int
	Paragraphs            []Paragraph
}

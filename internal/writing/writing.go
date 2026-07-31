package writing

type Task1Type string

type Task2Category string

const (
	Task1TypeLine        Task1Type = "Line"
	Task1TypeBar         Task1Type = "Bar"
	Task1TypePie         Task1Type = "Pie"
	Task1TypeTable       Task1Type = "Table"
	Task1TypeCombination Task1Type = "Combination"
	Task1TypeProcess     Task1Type = "Process"
	Task1TypeMap         Task1Type = "Map"
)

var AllTask1Types = []Task1Type{
	Task1TypeLine,
	Task1TypeBar,
	Task1TypePie,
	Task1TypeTable,
	Task1TypeCombination,
	Task1TypeProcess,
	Task1TypeMap,
}

func IsValidTask1Type(s string) bool {
	for _, t := range AllTask1Types {
		if string(t) == s {
			return true
		}
	}
	return false
}

const (
	Task2CategoryOpinion                    Task2Category = "Opinion"
	Task2CategoryDiscussion                 Task2Category = "Discussion"
	Task2CategoryProblemAndSolution         Task2Category = "Problem And Solution"
	Task2CategoryAdvantagesAndDisadvantages Task2Category = "Advantages And Disadvantages"
	Task2CategoryTwoPartQuestion            Task2Category = "Two-Part Question"
)

var AllTask2Categories = []Task2Category{
	Task2CategoryOpinion,
	Task2CategoryDiscussion,
	Task2CategoryProblemAndSolution,
	Task2CategoryAdvantagesAndDisadvantages,
	Task2CategoryTwoPartQuestion,
}

func IsValidTask2Category(s string) bool {
	for _, c := range AllTask2Categories {
		if string(c) == s {
			return true
		}
	}
	return false
}

type Paragraph struct {
	Text string
}

type Task1 struct {
	Id       string    `json:"id"`
	Question string    `json:"question"`
	Type     Task1Type `json:"type"`
	ImageKey string    `json:"image_key"`
}

type Task2 struct {
	Id       string        `json:"id"`
	Question string        `json:"question"`
	Category Task2Category `json:"category"`
}

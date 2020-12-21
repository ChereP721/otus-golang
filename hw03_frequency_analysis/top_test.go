package hw03_frequency_analysis //nolint:golint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

type test struct {
	name     string
	input    string
	expected []string
}

var testList = [...]test{
	{
		name:     "frequency words < 10",
		input:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		expected: []string{"in", "ut", "dolor", "dolore"},
	},
	{
		name:     "no frequency words",
		input:    "Lorem ipsum sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt labore et magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi aliquip ex ea commodo consequat. Duis aute irure dolor reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt culpa qui officia deserunt mollit anim id est laborum.",
		expected: []string{},
	},
	{
		name:     "specialchars",
		input:    "###one### two three ^^ ^^ $$ !!!one!!! ---one--- two",
		expected: []string{"one", "two"},
	},
	{
		name:     "specialchars in word",
		input:    "###one### two~~~three !!!one!!! ---one--- two",
		expected: []string{"one", "two"},
	},
	{
		name:     "digits in word",
		input:    "123one456 two789three 123one456 789one000 two",
		expected: []string{"one", "two"},
	},
	{
		name:     "russian yo",
		input:    "Для использования буквы \"ё\" регулярку пришлось слегка дописать, так как ё находится после буквы я",
		expected: []string{"ё", "буквы"},
	},
	{
		name:     "slashes and regexp special symbols",
		input:    `\ \s ^ $ $ \p \d [] . * \+ \- \\ \n \r \t \n \r \t \`,
		expected: []string{"n", "r", "t"},
	},
	{
		name:     "other alphabets",
		input:    `Schulabgänger? Student? Absolvent? Berufserfahrener? Њ tessssstttt Њ šo 勝利 الرفاق 勝利 الرفاق Schulabgänger? Student? Absolvent? Berufserfahrener?`,
		expected: []string{"schulabgänger", "student", "absolvent", "berufserfahrener", "њ", "勝利", "الرفاق"},
	},
}

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			require.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			require.ElementsMatch(t, expected, Top10(text))
		}
	})

	for _, testCase := range testList {
		t.Run(testCase.name, func(t *testing.T) {
			require.ElementsMatch(t, testCase.expected, Top10(testCase.input))
		})
	}
}

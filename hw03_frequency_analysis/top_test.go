package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

var text2 = `Lorem ipsum dolor sit, consectetur adipiscing elit.
	Sed vel eros eget libero auctor dapibus. Vivamus bibendum eget sapien
	at bibendum. Morbi ligula elit, suscipit nunc ut, rutrum tempor dui.
	Maecenas vitae congue diam, ut tincidunt quam. Praesent tempus nisl, 
	vitae euismod erat ullamcorper vitae. Cras aliquam tellus quam, at elementum odio vulputate eu.
	Phasellus ut dolor a nulla malesuada laoreet ac at lorem.
	Etiam tempus, turpis nec sodales luctus, justo augue elementum lacus, vel 
	aliquam nisi ipsum in diam. Nunc ultrices magna id sagittis pharetra. Nulla 
	facilities. Ut nulla urna, auctor sit amet augue id, molestie finibus justo. 
	Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur sed lacinia magna, 
	at lobortis justo. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; 
	Maecenas leo eros, feugiat ac nisi sed, ullamcorper tristique neque. Donec euismod quam a auctor dictum. Aliquam
	nec lorem elementum, hendrerit ligula ac, tincidunt turpis. Cras dignissim ex ex. Suspendisse a mi ac 
	risus faucibus finibus dictum eu neque. Suspendisse tempor, tellus vel sagittis varius, massa magna aliquam 
	erat, ut porttitor neque libero a ipsum. Nunc vel fermentum tortor. Suspendisse luctus dolor nibh.
	Fusce molestie condimentum purus vel feugiat. Curabitur sollicitudin metus sit ipsum congue, 
	sed porta ipsum gravida. Nullam nisl ante, consectetur eu pellentesque et, efficitur ac. 
	Fusce pretium tortor, eget cursus sem. In malesuada turpis eu turpis pellentesque, sit amet aliquet 
	sem tincidunt. Nullam a turpis arcu. Nam suscipit nisl id tortor suscipit molestie.
	Sed tristique commodo sapien. Donec ac iaculis libero, nec bibendum ex. Quisque ultricies est 
	mauris, in commodo sem ultrices. Donec nisl porta, iaculis mauris a, tempus nunc. Sed 
	finibus pellentesque risus a vulputate. Nulla ornare aliquet leo in accumsan. Fusce aliquam libero ut 
	ex placerat tempus. Cras ut elit tellus. Suspendisse id urna urna. Fusce luctus nec nunc et convallis. 
	Mauris vitae lobortis magna. Nam feugiat, diam eget dapibus hendrerit, ex nunc sollicitudin nisi, at
	consectetur quam ante ac erat. Mauris rhoncus vestibulum. Nunc pretium nunc sit amet 
	viverra scelerisque. Etiam efficitur odio a sapien porttitor, eu egestas lacus tristique. 
	Proin aliquam nisl sit amet feugiat egestas.
	Donec dapibus justo quis consequat lacinia. Morbi non ornare dolor. Mauris velit quam,
	euismod ut consequat at, maximus et tellus. Proin faucibus auctor ligula, non consectetur
	nibh ultrices quis. Etiam porta lacus et nunc scelerisque, a tincidunt magna condimentum.
	Praesent imperdiet convallis imperdiet. Quisque consequat sodales orci, sed semper nunc auctor at.`

func TestTop10(t *testing.T) {
	t.Parallel()

	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	tt := []struct {
		name   string
		text   string
		exp    []string
		notExp []string
	}{
		{
			name: "vinny",
			text: text,
			exp: []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			},
			notExp: []string{ // Если выполнено задание со звездочкой не валидный тест
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			},
		},
		{
			text: text2,
			name: "lorem",
			exp: []string{
				"nunc",    // 10
				"a",       // 9
				"ut",      // 8
				"ac",      // 7
				"at",      // 7
				"ipsum",   // 7
				"sed",     // 7
				"sit",     // 7
				"aliquam", // 6
				"amet",    // 5
			},
			notExp: []string{
				"aliquam", // 10
				"nunc",    // 9
				"ut",      // 8
				"ac",      // 7
				"at",      // 7
				"ipsum",   // 6
				"amet",    // 6
				"sit",     // 6
				"aliquam", // 6
				"sed",     // 6
			},
		},
	}

	t.Run("positive test", func(t *testing.T) {
		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				require.Equal(t, tc.exp, Top10(tc.text))
			})
		}
	})

	t.Run("negative test", func(t *testing.T) {
		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				require.NotEqual(t, tc.notExp, Top10(tc.text))
			})
		}
	})
}

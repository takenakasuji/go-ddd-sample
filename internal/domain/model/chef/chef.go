package chef

import (
	"github.com/takenakasuji/go-ddd-sample/internal/dto"
)

type Chef struct {
	name        *Name
	catchphrase string
	description string
}

// NewChef entityなので識別子も引数にとるべきだがRDBの識別子が存在しないので悩み中
func NewChef(name *Name) *Chef {
	chef := new(Chef)
	chef.name = name
	return chef
}

// ChangeCatchphrase Chefエンティティのcatchphraseは可変性があることを仮定し変更のメソッドを持つ
func (chef *Chef) ChangeCatchphrase(catchphrase string) {
	//	文字数制限などドメインの成約を記述
	chef.catchphrase = catchphrase
}

// ChangeDescription Chefエンティティのdescriptionは可変性があることを仮定し変更のメソッドを持つ
func (chef *Chef) ChangeDescription(description string) {
	//	文字数制限などドメインの成約を記述
	chef.description = description
}

// Equals 等価性の比較、エンティティであるため識別子によって比較される
// 今回は識別子の定義ができてないのでコメントとして参考実装を書いておく
//func (id *ID) Equals(otherID ID) bool {
//	return reflect.DeepEqual(id, otherID)
//}

func (chef *Chef) ToDto() dto.Chef {
	return dto.Chef{
		Name:        chef.name.value,
		Catchphrase: chef.catchphrase,
		Description: chef.description,
	}
}

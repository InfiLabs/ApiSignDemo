/**
 * @author  zhaoliang.liang
 * @date  2024/1/19 0019 11:20
 */

package define

type InfiUserType string

const (
	InfiUserTypeEditor  InfiUserType = "editor"  // 编辑者
	InfiUserTypeOwner   InfiUserType = "owner"   // 所有者
	InfiUserTypeVisitor InfiUserType = "visitor" // 访客
)

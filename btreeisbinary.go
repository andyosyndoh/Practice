package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Data < root.Left.Data {
		return false
	}
	if root.Right != nil && root.Data > root.Right.Data {
		return false
	}
	return true
}

package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right.Data > root.Data {
		return root.Right
	}
	if root.Left.Data > root.Data {
		return root.Left
	}
	return root.Right
}

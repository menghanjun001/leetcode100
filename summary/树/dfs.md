- dfs模板

```java
Set<Node> visited = new HashSet<>();

public void dfs(Node root, Set<Node> visited) {
    if (visited.contains(root)) {//terminator
        // already visited
        return;
    }
    visited.add(root);

    // process current node here.
    for (Node nextNode : node.children) {
        if (!visited.contains(nextNode)) {
            dfs(nextNode, visited);
        }
    }
}
```

- 39.组合总和
- 46.全排列
- 78.子集
- 79.单词搜索
- 124.二叉树中的最大路径和


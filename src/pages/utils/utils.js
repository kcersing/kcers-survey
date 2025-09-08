"use strict";
function treeToArray(treeNodes) {
    let result = [];
    //递归函数 traverse，用于处理单个节点
    function traverse(node) {
        const newNode = { ...node };
        delete newNode.children;
        // 将没有子节点的新节点添加到结果数组中
        result.push(newNode);
        // 如果当前节点包含 children 属性（即有子节点）
        if (node.children) {
            node.children.forEach(traverse);
        }
    }
    treeNodes.forEach(traverse);
    return result;
}
//# sourceMappingURL=utils.js.map
/**
 * 共享跳题规则处理 — 所有题目组件的 onChange 统一使用此函数。
 * 选中时记录跳转目标，点"下一步"时才真正跳转。
 */
export function handleJump(
  question: API.Questions,
  value: string | number | string[],
  setCurrent: (step: number) => void,
) {
  if (!question.jumpRules?.length) {
    setCurrent(-1); // 清除旧跳转目标
    return;
  }

  for (const rule of question.jumpRules) {
    if (rule.operators !== 'equals') continue;

    const matches = Array.isArray(value)
      ? value.includes(rule.answer)
      : String(value) === rule.answer;

    if (matches) {
      setCurrent(parseInt(rule.nextQuestionId));
      return;
    }
  }

  setCurrent(-1); // 无匹配，清除旧跳转目标
}

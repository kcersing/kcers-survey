/**
 * 共享跳题规则处理 — 所有题目组件的 onChange 统一使用此函数。
 *
 * @param question   当前题目
 * @param value      用户选择的值（单选 string、多选 string[]、评分 number）
 * @param setCurrent StepsForm 的 current setter
 */
export function handleJump(
  question: API.Questions,
  value: string | number | string[],
  setCurrent: (step: number) => void,
) {
  if (!question.jumpRules?.length) return;

  for (const rule of question.jumpRules) {
    if (rule.operators !== 'equals') continue;

    const matches = Array.isArray(value)
      ? value.includes(rule.answer)
      : String(value) === rule.answer;

    if (matches) {
      setCurrent(parseInt(rule.nextQuestionId));
    }
  }
}

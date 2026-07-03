import { useRef, useCallback } from 'react';

interface AnswerParams {
  surveyId: number;
  questionId: number;
  type: string;
  sn: string;
}

type SubmitFn = (fields: Record<string, any>) => Promise<boolean>;

/**
 * 统一答题提交 hook，支持 debounce。
 *
 * - delay = 0: 即时提交（选择类组件）
 * - delay > 0: debounce 提交（文本/数字输入组件）
 */
export function useDebouncedAnswer(
  params: AnswerParams,
  addRespondent: SubmitFn,
  delay: number = 0,
) {
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const latestRef = useRef<any>(null);
  const submittingRef = useRef(false);

  const flush = useCallback(async () => {
    const value = latestRef.current;
    if (!value || submittingRef.current) return;
    submittingRef.current = true;
    try {
      await addRespondent({
        surveyId: params.surveyId,
        questionId: params.questionId,
        type: params.type,
        value,
        sn: params.sn,
      });
    } finally {
      submittingRef.current = false;
    }
  }, [params, addRespondent]);

  const submitAnswer = useCallback(
    (value: string[]) => {
      latestRef.current = value;

      if (delay <= 0) {
        // 即时提交
        if (timerRef.current) {
          clearTimeout(timerRef.current);
          timerRef.current = null;
        }
        addRespondent({
          surveyId: params.surveyId,
          questionId: params.questionId,
          type: params.type,
          value,
          sn: params.sn,
        });
      } else {
        // debounce
        if (timerRef.current) clearTimeout(timerRef.current);
        timerRef.current = setTimeout(() => {
          flush();
        }, delay);
      }
    },
    [params, addRespondent, delay, flush],
  );

  return { submitAnswer, flush };
}

export default useDebouncedAnswer;

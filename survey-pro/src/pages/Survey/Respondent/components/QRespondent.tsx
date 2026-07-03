import React, { useEffect, useRef } from 'react';
import { ProFormText, StepsForm } from '@ant-design/pro-components';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';

const DEBOUNCE_MS = 400;

const QRespondent = (props: QuestionComponentProps) => {
  const { surveyId, questions, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const timers = useRef<Record<string, ReturnType<typeof setTimeout>>>({});

  useEffect(() => {
    setCurrentNum(0);
  }, []);

  const debouncedAdd = (type: string, val: string) => {
    if (timers.current[type]) clearTimeout(timers.current[type]);
    timers.current[type] = setTimeout(() => {
      addRespondent({
        surveyId,
        questionId: 0,
        type,
        value: [val],
        sn: generateRandom,
      });
    }, DEBOUNCE_MS);
  };

  return (
    <StepsForm.StepForm
      name={`key_${questions.length + 1}`}
      key={`key_${questions.length + 1}`}
    >
      <ProFormText
        width="md"
        onChange={(e) => debouncedAdd('respondent', e.target.value)}
        label="访谈人姓名"
        rules={[{ required: true }]}
        name="respondent"
      />
      <ProFormText
        width="md"
        onChange={(e) => debouncedAdd('respondentPhone', e.target.value)}
        label="联系电话"
        rules={[{ required: true, len: 11 }]}
        name="respondentPhone"
      />
      <ProFormText
        width="md"
        onChange={(e) => debouncedAdd('researcher', e.target.value)}
        label="调研员姓名"
        rules={[{ required: true }]}
        name="researcher"
      />
      <ProFormText
        width="md"
        onChange={(e) => debouncedAdd('researcherPhone', e.target.value)}
        label="联系电话"
        rules={[{ required: true, len: 11 }]}
        name="researcherPhone"
      />
    </StepsForm.StepForm>
  );
};

export default QRespondent;

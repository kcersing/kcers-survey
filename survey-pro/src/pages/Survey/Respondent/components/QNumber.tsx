import React, { useState, useRef } from 'react';
import { Form } from 'antd';
import { ProFormDigit } from '@ant-design/pro-components';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QNumber = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState<number | null>(null);
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  if (!question) return null;

  const onChange = (e: number | null) => {
    if (e == null) return;
    setValue(e);

    // debounce 500ms — 数字输入稍长 debounce
    if (timerRef.current) clearTimeout(timerRef.current);
    timerRef.current = setTimeout(() => {
      addRespondent({
        surveyId,
        questionId: question.id,
        type: question.type,
        value: [e.toString()],
        sn: generateRandom,
      });
    }, 500);
  };

  const onBlur = () => {
    if (timerRef.current) clearTimeout(timerRef.current);
    if (value != null) {
      addRespondent({
        surveyId,
        questionId: question.id,
        type: question.type,
        value: [value.toString()],
        sn: generateRandom,
      });
    }
    handleJump(question, value ?? '', setCurrent);
  };

  return (
    <Form.Item name={['question', "'" + question.id + "'"]}>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <ProFormDigit
        width="md"
        placeholder="请输入..."
        name={['question', question.id]}
        style={{ width: 60 }}
        onChange={onChange}
        onBlur={onBlur}
        rules={[{ required: question.required === 1, message: '必填项' }]}
      />
      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
        value={value}
      />
    </Form.Item>
  );
};

export default QNumber;

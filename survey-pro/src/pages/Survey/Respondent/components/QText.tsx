import React, { useState, useRef } from 'react';
import { Form } from 'antd';
import { ProFormTextArea } from '@ant-design/pro-components';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QText = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState('');
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  if (!question) return null;

  const onChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const val = e.target.value;
    setValue(val);

    // debounce 300ms — 输入停止后才提交
    if (timerRef.current) clearTimeout(timerRef.current);
    timerRef.current = setTimeout(() => {
      addRespondent({
        surveyId,
        type: question.type,
        questionId: question.id,
        value: [val],
        sn: generateRandom,
      });
    }, 300);
  };

  const onBlur = () => {
    // 失焦时立即提交最新值
    if (timerRef.current) clearTimeout(timerRef.current);
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [value],
      sn: generateRandom,
    });
    handleJump(question, value, setCurrent);
  };

  return (
    <Form.Item name={['question', "'" + question.id + "'"]}>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <ProFormTextArea
        width="md"
        name={['question', question.id]}
        onChange={onChange}
        onBlur={onBlur}
        placeholder={question.remark}
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

export default QText;

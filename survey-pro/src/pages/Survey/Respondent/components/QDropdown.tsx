import React, { useState } from 'react';
import { Form, Select } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QDropdown = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState<string>('');
  if (!question) return null;

  const onChange = (val: string) => {
    setValue(val);
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [val],
      sn: generateRandom,
    });
    handleJump(question, val, setCurrent);
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请选择一个选项' }]}
      >
        <Select
          placeholder="请选择..."
          onChange={onChange}
          style={{ width: '100%', maxWidth: 400 }}
          options={question.options?.map((o) => ({
            label: o.content,
            value: o.content,
          })) ?? []}
        />
      </Form.Item>
      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
        value={value}
      />
    </>
  );
};

export default QDropdown;

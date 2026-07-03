import React, { useState } from 'react';
import { Button, Form } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const scores = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

const QNps = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState<number | null>(null);
  if (!question) return null;

  const onChange = (val: number) => {
    setValue(val);
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [val.toString()],
      sn: generateRandom,
    });
    handleJump(question, val, setCurrent);
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请选择评分' }]}
      >
        <div style={{ display: 'flex', justifyContent: 'space-between', maxWidth: 480, marginBottom: 8 }}>
          {scores.map((n) => (
            <Button
              key={n}
              type={value === n ? 'primary' : 'default'}
              shape="circle"
              size="small"
              onClick={() => onChange(n)}
            >
              {n}
            </Button>
          ))}
        </div>
        <div style={{ display: 'flex', justifyContent: 'space-between', maxWidth: 480, color: '#999', fontSize: 12 }}>
          <span>完全不可能</span>
          <span>非常可能</span>
        </div>
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

export default QNps;

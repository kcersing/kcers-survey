import React, { useState } from 'react';
import { Form } from 'antd';
import { ProFormRate } from '@ant-design/pro-components';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const rateCharacter = ({ index = 0 }: { index?: number }) => index + 1;

const QRate = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState(0);
  if (!question) return null;

  const onChange = (e: number) => {
    addRespondent({
      surveyId,
      questionId: question.id,
      type: question.type,
      value: [e.toString()],
      sn: generateRandom,
    });
    handleJump(question, e, setCurrent);
  };

  const getCharacter = question.options?.[0]?.serial
    ? ({ index = 0 }: { index?: number }) => index + question.options[0].serial
    : rateCharacter;

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item name={['question', "'" + question.id + "'"]}>
        <ProFormRate
          style={{ color: 'rgba(150, 205, 205, 0.6)' }}
          fieldProps={{
            character: getCharacter,
            allowHalf: false,
            count: question.options?.[0]?.inputs ?? 5,
            style: { fontSize: '30px' },
          }}
          name={['question', question.id]}
          onChange={onChange}
          rules={[{ required: question.required === 1, message: '必填项' }]}
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

export default QRate;

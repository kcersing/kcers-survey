import React, { useState, useRef } from 'react';
import type { RadioChangeEvent } from 'antd';
import { Input, Form } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import { ProFormRadio } from '@ant-design/pro-components';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const SingleChoice = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState('');
  if (!question) return null;

  const onChange = (e: RadioChangeEvent) => {
    setValue(e.target.value);
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [e.target.value.toString()],
      sn: generateRandom,
    });
    handleJump(question, e.target.value, setCurrent);
  };

  const otherTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const onOtherInput = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const val = e.target.value.toString();
    if (otherTimerRef.current) clearTimeout(otherTimerRef.current);
    otherTimerRef.current = setTimeout(() => {
      addRespondent({
        surveyId,
        type: 'input',
        questionId: question.id,
        value: [val],
        sn: generateRandom,
      });
    }, 300);
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '这是必填项' }]}
      >
        <ProFormRadio.Group
          onChange={onChange}
          style={style}
          layout="vertical"
          fieldProps={{ style: { fontSize: '2em' } }}
          options={question.options.map((option) => ({
            value: option.content,
            label: option.inputs !== 2 ? option.content : (
              <>
                {option.content}...
                {value === option.content && (
                  <Input
                    onChange={onOtherInput}
                    variant="filled"
                    placeholder="请输入..."
                    style={{ width: 120, marginInlineStart: 12 }}
                  />
                )}
              </>
            ),
          }))}
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

export default SingleChoice;

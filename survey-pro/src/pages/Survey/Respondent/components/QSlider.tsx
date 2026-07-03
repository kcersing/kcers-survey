import React, { useState, useRef } from 'react';
import { Form, Slider } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QSlider = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const min = question.options?.[0]?.serial ?? 0;
  const max = question.options?.[0]?.inputs ?? 100;
  const [value, setValue] = useState<number>(min);
  const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  if (!question) return null;

  const onChange = (val: number) => {
    setValue(val);
    if (timerRef.current) clearTimeout(timerRef.current);
    timerRef.current = setTimeout(() => {
      addRespondent({
        surveyId,
        type: question.type,
        questionId: question.id,
        value: [val.toString()],
        sn: generateRandom,
      });
    }, 200);
  };

  const onAfterChange = (val: number) => {
    // 拖动结束后立即提交
    if (timerRef.current) clearTimeout(timerRef.current);
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [val.toString()],
      sn: generateRandom,
    });
    handleJump(question, val, setCurrent);
  };

  const marks: Record<number, string> = {
    [min]: String(min),
    [max]: String(max),
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '必填项' }]}
      >
        <Slider
          min={min} max={max} marks={marks}
          onChange={onChange}
          onAfterChange={onAfterChange}
          style={{ maxWidth: 400, marginBottom: 8 }}
          tooltip={{ formatter: (v) => String(v) }}
        />
        <div style={{ color: '#888' }}>当前值: {value}</div>
      </Form.Item>
      <QJumpRules
        surveyId={surveyId} question={question} generateRandom={generateRandom}
        addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}
        value={value}
      />
    </>
  );
};

export default QSlider;

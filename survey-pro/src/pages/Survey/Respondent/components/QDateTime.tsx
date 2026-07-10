import React, { useState } from 'react';
import { DatePicker, Form } from 'antd';
import dayjs from 'dayjs';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QDateTime = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState('');
  const form = Form.useFormInstance();
  if (!question) return null;

  const fieldName = ['question', "'" + question.id + "'"];

  return (
    <Form.Item
      name={fieldName}
      rules={[{ required: question.required === 1, message: '必填项' }]}
    >
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <DatePicker
        showTime
        placeholder="请选择日期和时间"
        format="YYYY-MM-DD HH:mm:ss"
        onChange={(date: dayjs.Dayjs | null) => {
          form.setFieldValue(fieldName, date);
          const formatted = date?.format('YYYY-MM-DD HH:mm:ss') || '';
          setValue(formatted);
          addRespondent({
            surveyId,
            questionId: question.id,
            type: question.type,
            value: [formatted],
            sn: generateRandom,
          });
          handleJump(question, formatted, setCurrent);
        }}
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

export default QDateTime;

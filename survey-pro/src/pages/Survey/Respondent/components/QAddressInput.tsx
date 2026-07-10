import React, { useRef } from 'react';
import { Form, Input } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';

const fields = [
  { name: 'area', label: '省（自治区、直辖市）', field: 'area' as const },
  { name: 'city', label: '市（自治州、地区、盟）', field: 'city' as const },
  { name: 'district', label: '县（自治县、县级市、区）', field: 'district' as const },
  { name: 'village', label: '镇（乡、街道）', field: 'village' as const },
  { name: 'address', label: '村', field: 'address' as const },
];

const QAddressInput = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const addrRef = useRef<Record<string, string>>({});
  if (!question) return null;

  const onChange = (field: string, val: string) => {
    addrRef.current = { ...addrRef.current, [field]: val };
    // 每个字段单独提交
    addRespondent({
      surveyId,
      type: field,
      questionId: question.id,
      value: [val],
      sn: generateRandom,
    });
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      {fields.map((f) => (
        <Form.Item key={f.name} label={f.label}>
          <Input
            placeholder={`请输入${f.label}`}
            onChange={(e) => onChange(f.field, e.target.value)}
          />
        </Form.Item>
      ))}
      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
        value={addrRef.current}
      />
    </>
  );
};

export default QAddressInput;
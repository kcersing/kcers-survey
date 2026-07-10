import React, { useState, useRef } from 'react';
import type { CheckboxValueType } from 'antd/es/checkbox/Group';
import { Input, Form } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import { ProFormCheckbox } from '@ant-design/pro-components';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const style: React.CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: 8,
};

const MultipleChoice = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [value, setValue] = useState<CheckboxValueType[]>([]);
  const [otherValue, setOtherValue] = useState('');
  const form = Form.useFormInstance();
  if (!question) return null;

  const maxSelect = question.valueNumber ? question.valueNumber : 999;

  const otherOption = question.options.find((o: any) => o.inputs === 2);
  const isOtherSelected = otherOption ? value.includes(otherOption.content) : false;
  const otherFieldName = ['question', 'other_' + question.id];

  const onChange = (checkedValues: CheckboxValueType[]) => {
    setValue(checkedValues);
    if (!checkedValues.includes(otherOption?.content || '')) {
      setOtherValue('');
      form.setFieldValue(otherFieldName, '');
    }
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: checkedValues,
      sn: generateRandom,
    });
    handleJump(question, checkedValues as string[], setCurrent);
  };

  const otherTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const onOtherInput = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const val = e.target.value.toString();
    setOtherValue(val);
    form.setFieldValue(otherFieldName, val);
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

  const isOptionDisabled = (content: string) => {
    if (value.length >= maxSelect) {
      return !value.includes(content);
    }
    return false;
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '这是必填项' }]}
      >
        <ProFormCheckbox.Group
          onChange={onChange}
          layout="vertical"
          style={style}
          options={question.options.map((option) => ({
            disabled: isOptionDisabled(option.content),
            value: option.content,
            label: option.inputs !== 2 ? option.content : (
              <>
                {option.content}...
                {value.length > 0 && value.includes(option.content) && (
                  <Input
                    value={otherValue}
                    onChange={onOtherInput}
                    variant="filled"
                    placeholder="请输入...(必填)"
                    status={isOtherSelected && !otherValue.trim() ? 'error' : undefined}
                    style={{ width: 120, marginInlineStart: 12 }}
                  />
                )}
              </>
            ),
          }))}
        />
      </Form.Item>
      {otherOption && (
        <Form.Item
          name={otherFieldName}
          rules={[{ required: isOtherSelected, message: '请填写其他内容' }]}
          noStyle
        >
          <Input style={{ display: 'none' }} />
        </Form.Item>
      )}
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

export default MultipleChoice;

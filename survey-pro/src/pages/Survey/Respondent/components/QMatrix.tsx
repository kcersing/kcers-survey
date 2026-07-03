import React, { useState } from 'react';
import { Form, Radio, Space, Button } from 'antd';
import type { RadioChangeEvent } from 'antd';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QMatrix = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [values, setValues] = useState<Record<number, string>>({});
  if (!question) return null;

  const rows = question.children ?? [];
  const cols = question.options ?? [];

  const onRowChange = (rowId: number, e: RadioChangeEvent) => {
    setValues((prev) => ({ ...prev, [rowId]: e.target.value }));
  };

  const submitMatrix = () => {
    const next = { ...values };
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [JSON.stringify(next)],
      sn: generateRandom,
    });
    handleJump(question, Object.values(next).join(','), setCurrent);
  };

  const allAnswered = rows.every((row) => values[row.id]);

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请完成矩阵' }]}
      >
        <div style={{ overflowX: 'auto' }}>
          <table style={{ borderCollapse: 'collapse', width: '100%', maxWidth: 600 }}>
            <thead>
              <tr>
                <th style={{ padding: 8 }}></th>
                {cols.map((col) => (
                  <th key={col.id} style={{ padding: 8, textAlign: 'center', fontSize: 13 }}>{col.content}</th>
                ))}
              </tr>
            </thead>
            <tbody>
              {rows.map((row) => (
                <tr key={row.id} style={{ borderTop: '1px solid #f0f0f0' }}>
                  <td style={{ padding: 8 }}>{row.content}</td>
                  <td colSpan={cols.length} style={{ padding: 4 }}>
                    <Radio.Group value={values[row.id]} onChange={(e) => onRowChange(row.id, e)}>
                      <Space>
                        {cols.map((col) => (
                          <Radio key={col.id} value={col.content}>{col.content}</Radio>
                        ))}
                      </Space>
                    </Radio.Group>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        <Button type="primary" onClick={submitMatrix} disabled={!allAnswered} style={{ marginTop: 12 }}>
          确认提交
        </Button>
      </Form.Item>
      <QJumpRules
        surveyId={surveyId} question={question} generateRandom={generateRandom}
        addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}
        value={Object.values(values)}
      />
    </>
  );
};

export default QMatrix;

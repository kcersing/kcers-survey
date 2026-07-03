import React, { useState, useCallback } from 'react';
import { Form, Button } from 'antd';
import { ArrowUpOutlined, ArrowDownOutlined } from '@ant-design/icons';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QRanking = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const [items, setItems] = useState(question?.options ?? []);
  if (!question) return null;

  const moveItem = useCallback((index: number, direction: -1 | 1) => {
    const next = [...items];
    const target = index + direction;
    if (target < 0 || target >= next.length) return;
    [next[index], next[target]] = [next[target], next[index]];
    setItems(next);
  }, [items]);

  const finish = () => {
    const rankValue = items.map((it) => it.content).join(' > ');
    addRespondent({
      surveyId,
      type: question.type,
      questionId: question.id,
      value: [rankValue],
      sn: generateRandom,
    });
    handleJump(question, rankValue, setCurrent);
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请排序' }]}
      >
        <div style={{ maxWidth: 400 }}>
          {items.map((item, index) => (
            <div
              key={item.id ?? index}
              style={{
                display: 'flex', alignItems: 'center', justifyContent: 'space-between',
                padding: '8px 12px', marginBottom: 4, background: '#fafafa',
                border: '1px solid #d9d9d9', borderRadius: 6,
              }}
            >
              <span>{index + 1}. {item.content}</span>
              <div>
                <Button size="small" icon={<ArrowUpOutlined />} disabled={index === 0} onClick={() => moveItem(index, -1)} />
                <Button size="small" icon={<ArrowDownOutlined />} disabled={index === items.length - 1} onClick={() => moveItem(index, 1)} style={{ marginLeft: 4 }} />
              </div>
            </div>
          ))}
          <Button type="primary" onClick={finish} style={{ marginTop: 12 }}>
            确认排序
          </Button>
        </div>
      </Form.Item>
      <QJumpRules
        surveyId={surveyId} question={question} generateRandom={generateRandom}
        addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}
        value={items.map((it) => it.content).join(' > ')}
      />
    </>
  );
};

export default QRanking;

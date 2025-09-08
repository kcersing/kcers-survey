import React, { useState } from 'react';
import { Form } from 'antd';
import { ProFormRate } from '@ant-design/pro-components';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
const QRate = (props) => {
    const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
    const [value, setValue] = useState(0);
    if (!question) {
        return null;
    }
    const onChange = (e) => {
        console.log(e);
        addRespondent({
            surveyId: surveyId,
            questionId: question.id,
            type: question.type,
            value: [e.toString()],
            sn: generateRandom,
        });
        if (question.jumpRules) {
            for (const jumpRule of question.jumpRules) {
                if (jumpRule.operators === 'equals' && String(e) === jumpRule.answer) {
                    // setCurrentNum(parseInt(jumpRule.nextQuestionId)-1);
                    setCurrent(parseInt(jumpRule.nextQuestionId));
                }
            }
        }
    };
    return (<Form.Item name={['question', "'" + question.id + "'"]} required={question.required === 1}>
      <h3>{question.serial ? question.serial + "-" : ""}{question.content}</h3>
    <ProFormRate name={['question', question.id]} onChange={onChange} rules={[{ required: question.required === 1, message: '必填项' }]}/>


      <QJumpRules surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent} value={value}/>
    </Form.Item>);
};
export default QRate;
//# sourceMappingURL=QRate.jsx.map
import React from 'react';
import SingleChoice from '@/pages/survey/respondent/components/SingleChoice';
import MultipleChoice from '@/pages/survey/respondent/components/MultipleChoice';
import QText from '@/pages/survey/respondent/components/QText';
import QNumber from '@/pages/survey/respondent/components/QNumber';
import QDate from '@/pages/survey/respondent/components/QDate';
import QRate from '@/pages/survey/respondent/components/QRate';
const QuestuinSun = (props) => {
    const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
    if (question.type === "single_choice") {
        return (<SingleChoice surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}/>);
    }
    if (question.type === 'multiple_choice') {
        return (<MultipleChoice surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}/>);
    }
    if (question.type === 'text') {
        return (<QText surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}/>);
    }
    if (question.type === 'number') {
        return (<QNumber surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}></QNumber>);
    }
    if (question.type === 'date') {
        return (<QDate surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}></QDate>);
    }
    if (question.type === 'rate') {
        return (<QRate surveyId={surveyId} question={question} generateRandom={generateRandom} addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}></QRate>);
    }
    return (<></>);
};
export default QuestuinSun;
//# sourceMappingURL=QuestuinSun.jsx.map
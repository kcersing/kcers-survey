import React from 'react';
import QuestuinSun from '@/pages/survey/respondent/components/QuestuinSun';
import type { QuestionWithValueProps } from '@/pages/survey/respondent/types';

const QJumpRules = (props: QuestionWithValueProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent, value } = props;

  if (!question.jumpRules?.length || !value) return null;

  if (typeof value === 'undefined' || (typeof value === 'string' && value.length === 0) || (Array.isArray(value) && value.length === 0)) {
    return null;
  }

  const matchedRules = question.jumpRules.filter((rule: any) => {
    if (rule.operators !== 'sub') return false;
    if (Array.isArray(value)) return value.includes(rule.answer);
    return String(value) === rule.answer;
  });

  if (!matchedRules.length) return null;

  const nextIds = new Set(matchedRules.map((r: any) => r.nextQuestionId));
  const matchedChildren = question.children?.filter((child: any) => nextIds.has(child.id)) ?? [];

  return (
    <>
      {matchedChildren.map((child: any) => (
        <QuestuinSun
          key={child.id}
          surveyId={surveyId}
          question={child}
          generateRandom={generateRandom}
          addRespondent={addRespondent}
          setCurrentNum={setCurrentNum}
          setCurrent={setCurrent}
        />
      ))}
    </>
  );
};

export default QJumpRules;

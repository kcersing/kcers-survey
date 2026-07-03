import React from 'react';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import SingleChoice from '@/pages/survey/respondent/components/SingleChoice';
import MultipleChoice from '@/pages/survey/respondent/components/MultipleChoice';
import QText from '@/pages/survey/respondent/components/QText';
import QNumber from '@/pages/survey/respondent/components/QNumber';
import QDate from '@/pages/survey/respondent/components/QDate';
import QRate from '@/pages/survey/respondent/components/QRate';
import QDateTime from '@/pages/survey/respondent/components/QDateTime';
import { QImage, QFile, QVideo, QAudio } from '@/pages/survey/respondent/components/QUpload';
import QPage from '@/pages/survey/respondent/components/QPage';
import QDropdown from '@/pages/survey/respondent/components/QDropdown';
import QSlider from '@/pages/survey/respondent/components/QSlider';
import QRanking from '@/pages/survey/respondent/components/QRanking';
import QMatrix from '@/pages/survey/respondent/components/QMatrix';
import QNps from '@/pages/survey/respondent/components/QNps';
import QSignature from '@/pages/survey/respondent/components/QSignature';

const componentMap: Record<string, React.ComponentType<QuestionComponentProps>> = {
  single_choice: SingleChoice,
  multiple_choice: MultipleChoice,
  text: QText,
  number: QNumber,
  date: QDate,
  rate: QRate,
  datetime: QDateTime,
  image: QImage,
  file: QFile,
  video: QVideo,
  audio: QAudio,
  page: QPage,
  dropdown: QDropdown,
  slider: QSlider,
  ranking: QRanking,
  matrix: QMatrix,
  nps: QNps,
  signature: QSignature,
};

const QuestuinSun = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;

  const Component = componentMap[question.type];
  if (Component) {
    return (
      <Component
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
      />
    );
  }

  return null;
};

export default QuestuinSun;

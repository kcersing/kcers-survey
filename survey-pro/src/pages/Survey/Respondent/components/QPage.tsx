import React from 'react';
import { Divider } from 'antd';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';

const QPage = (props: QuestionComponentProps) => {
  const { question } = props;
  if (!question) return null;

  return (
    <Divider orientation="left" plain>
      {question.content || '下一页'}
    </Divider>
  );
};

export default QPage;

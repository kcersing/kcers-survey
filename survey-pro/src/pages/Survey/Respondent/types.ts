import type { ProFormInstance } from '@ant-design/pro-components';

export interface QuestionComponentProps {
  surveyId: number;
  question: API.Questions;
  generateRandom: string;
  addRespondent: (fields: Record<string, any>) => Promise<boolean>;
  setCurrentNum: (num: number) => void;
  setCurrent: (current: number) => void;
}

export interface QuestionWithValueProps extends QuestionComponentProps {
  value?: any;
}

export interface RespondentPageState {
  survey: API.Survey;
  questions: API.Questions[];
  current: number;
  currentNum: number;
  generateRandom: string;
  loading: boolean;
  latitude: number | null;
  longitude: number | null;
}

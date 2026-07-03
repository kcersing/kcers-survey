import React from 'react';
import { Form, message } from 'antd';
import { ProFormUploadButton } from '@ant-design/pro-components';
import { pubUpload } from '@/services/ant-design-pro/api';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';

interface UploadProps extends QuestionComponentProps {
  accept: string;
  listType?: 'picture-card' | 'picture' | 'text';
  maxCount?: number;
}

const QUpload = (props: UploadProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent, accept, listType = 'picture-card', maxCount = 1 } = props;
  if (!question) return null;

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请上传文件' }]}
      >
        <ProFormUploadButton
          name={['question', question.id]}
          label="点击上传"
          max={maxCount}
          fieldProps={{
            accept,
            listType,
            maxCount,
            customRequest: (options: any) => {
              const { file, onSuccess, onError } = options;
              pubUpload({ file }).then((res: any) => {
                if (res.code === 0) {
                  message.success('上传成功');
                  addRespondent({
                    surveyId,
                    type: question.type,
                    questionId: question.id,
                    value: [res.data.url],
                    sn: generateRandom,
                  });
                  onSuccess(res, file);
                } else {
                  message.error('上传失败');
                  onError(new Error('upload failed'));
                }
              }).catch((err: any) => {
                message.error('上传失败');
                onError(err);
              });
            },
          }}
        />
      </Form.Item>
      <QJumpRules
        surveyId={surveyId}
        question={question}
        generateRandom={generateRandom}
        addRespondent={addRespondent}
        setCurrentNum={setCurrentNum}
        setCurrent={setCurrent}
      />
    </>
  );
};

// 各媒体类型的薄封装
export const QImage = (props: QuestionComponentProps) => (
  <QUpload {...props} accept="image/*" listType="picture-card" maxCount={10} />
);

export const QFile = (props: QuestionComponentProps) => (
  <QUpload {...props} accept="*/*" listType="text" maxCount={1} />
);

export const QVideo = (props: QuestionComponentProps) => (
  <QUpload {...props} accept="video/*" listType="text" maxCount={1} />
);

export const QAudio = (props: QuestionComponentProps) => (
  <QUpload {...props} accept="audio/*" listType="text" maxCount={1} />
);

export default QUpload;

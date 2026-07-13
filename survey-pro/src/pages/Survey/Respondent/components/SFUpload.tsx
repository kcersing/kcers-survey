import React from 'react';
import { ProFormUploadButton } from '@ant-design/pro-components';
import { pubUpload } from '@/services/ant-design-pro/api';
import { message } from 'antd';

const SFUpload = (props: any) => {
  const { surveyId, generateRandom, addRespondent } = props;

  return (
    <ProFormUploadButton
      name="上传合照"
      label="上传合照"
      rules={[{ required: true, message: '请上传合照' }]}
      max={10}
      action={(file: any) => {
        pubUpload({ file }).then((res: any) => {
          if (res.code === 0) {
            message.success('上传成功');
            addRespondent({
              surveyId,
              type: 'image',
              value: [res.data.url],
              sn: generateRandom,
            });
          }
        });
      }}
      listType="picture-card"
      accept="image/*"
      maxSize={5 * 1024}
    />
  );
};

export default SFUpload;
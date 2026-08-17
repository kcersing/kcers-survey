import React from 'react';
import { ProFormUploadButton } from '@ant-design/pro-components';
import { pubUpload } from '@/services/ant-design-pro/api';
import { compressImage } from '@/utils/imageCompress';
import { message } from 'antd';

const SFUpload = (props: any) => {
  const { surveyId, generateRandom, addRespondent } = props;

  return (
    <ProFormUploadButton
      name="上传合照"
      label="上传合照"
      rules={[{ required: true, message: '请上传合照' }]}
      max={10}
      fieldProps={{
        accept: 'image/*',
        listType: 'picture-card',
        maxCount: 10,
        customRequest: (options: any) => {
          const { file, onSuccess, onError } = options;
          compressImage(file).then((compressed) => {
            pubUpload({ file: compressed }).then((res: any) => {
              if (res.retcode === 0) {
                message.success('上传成功');
                addRespondent({
                  surveyId,
                  type: 'image',
                  value: [res.url],
                  sn: generateRandom,
                });
                onSuccess(res, compressed);
              } else {
                message.error(res.retmsg || '上传失败');
                onError(new Error(res.retmsg || 'upload failed'));
              }
            }).catch((err: any) => {
              message.error('上传失败');
              onError(err);
            });
          }).catch((err: any) => {
            message.error('压缩失败');
            onError(err);
          });
        },
      }}
    />
  );
};

export default SFUpload;
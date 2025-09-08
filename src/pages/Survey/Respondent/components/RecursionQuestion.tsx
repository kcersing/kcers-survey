import React, {memo} from "react";
import {
  ProForm,
  ProFormCheckbox,
  ProFormDatePicker,
  ProFormDigit,
  ProFormRadio, ProFormRate,
  ProFormTextArea, ProFormUploadButton
} from "@ant-design/pro-components";

import API from "@/services/typings";

const QuestionRenderer = memo(({
                                 question,
                                 depth,
                                 form,
                                 answers,
                                 onUploadStart,
                                 onUploadSuccess,
                                 onUploadError
                               }) => {
  // 普通问题
  return (
    <div className={`question question-depth-${depth}`}>
      <ProForm.Item
        name={`question${question.id}`}
        label={question.content}
        rules={question.required===1 ? [{ required: true, message: '此字段为必填项' }] : []}
        initialValue={answers[question.id]?.[`question${question.id}`]}
      >
        {renderQuestionControl(question, depth, {
          onUploadStart,
          onUploadSuccess,
          onUploadError
        })}
      </ProForm.Item>
    </div>
  );
}, (prevProps, nextProps) => {
  // 浅比较问题和答案，避免不必要的渲染
  return prevProps.question.id === nextProps.question.id &&
    prevProps.answers[prevProps.question.id] === nextProps.answers[nextProps.question.id];
});

// 问题控件渲染
const renderQuestionControl = (
  question:API.Questions,
  depth: number,
  uploadHandlers: {
    onUploadStart: (type: string) => void;
    onUploadSuccess: (type: string, file: any) => void;
    onUploadError: (type: string, error: any) => void;
  }
) => {
  const { onUploadStart, onUploadSuccess, onUploadError } = uploadHandlers;
  console.log(question)

  switch (question.type) {
    case 'single_choice':
      return (
        <ProFormRadio.Group
          name={question.id}
          options={question.options.map(option => ({
            value:option.content,
            label: option.content,
          }))}
        />
      );

    case 'multiple_choice':
      return (
        <ProFormCheckbox.Group
          name={question.id}
          options={question.options.map(option => ({
            value:option.content,
            label: option.content,
          }))}
        />
      );

    case 'text':
      return <ProFormTextArea name={question.id} label="名称"  placeholder="请输入内容..." />;

    case 'number':
      return <ProFormDigit placeholder="请输入数字"
                           name={question.id}
                           min={1}
                           max={10} />;

    case 'date':
      return <ProFormDatePicker  name={question.id} placeholder="请选择日期" />;

    case 'rate':
      return (
        <ProFormRate name={question.id} label="Rate" />
      );

    case 'uploadImage':
      return (
        <ProFormUploadButton  name={question.id} label="Upload" />
        // <ProFormUploadDragger
        //   name="file"
        //   className="upload-image"
        //   onStart={() => onUploadStart('image')}
        //   onSuccess={(info) => onUploadSuccess('image', info)}
        //   onError={(error) => onUploadError('image', error)}
        // >
        //   <Button icon={<PlusOutlined />}>上传图片</Button>
        // </ProFormUploadDragger>
      );

    case 'uploadFile':
      return (
        <ProFormUploadButton name={question.id} label="Upload" />
        // <ProFormUploadDragger
        //   name="file"
        //   className="upload-file"
        //   onStart={() => onUploadStart('file')}
        //   onSuccess={(info) => onUploadSuccess('file', info)}
        //   onError={(error) => onUploadError('file', error)}
        // >
        //   <Button icon={<PlusOutlined />}>上传文件</Button>
        // </ProFormUploadDragger>
      );
    case 'h2':
      return (
        <h3>{question.content}</h3>
      );
    case 'h3':
      return (
        <h4>{question.content}</h4>
      );

    default:
      return null;
  }
};

const RecursionQuestion =(currentQuestion,currentPath,form,answers,handleUploadStart,handleUploadSuccess,handleUploadError)=>{
  return (
    <>
      {currentQuestion && (
        <QuestionRenderer
          question={currentQuestion}
          depth={currentPath.length}
          form={form}
          answers={answers}
          onUploadStart={handleUploadStart}
          onUploadSuccess={handleUploadSuccess}
          onUploadError={handleUploadError}
        />
      )}

      {/* 渲染子问题 - 只在有子问题时显示 */}
      {currentQuestion && currentQuestion.children && currentQuestion.children.length > 0 && (
        <div className="children-questions">
          {currentQuestion.children.map((child, index) => (
            <div key={child.id} style={{ paddingLeft: '20px' }}>

              { RecursionQuestion(child,currentPath,form,answers,handleUploadStart,handleUploadSuccess,handleUploadError)}

            </div>
          ))}
        </div>
      )}
    </>);
}

export default RecursionQuestion;

import React, { useRef, useState, useCallback } from 'react';
import { Form, Button, message } from 'antd';
import { ClearOutlined } from '@ant-design/icons';
import QJumpRules from '@/pages/survey/respondent/components/QJumpRules';
import type { QuestionComponentProps } from '@/pages/survey/respondent/types';
import { handleJump } from './jumpRules';

const QSignature = (props: QuestionComponentProps) => {
  const { surveyId, question, generateRandom, addRespondent, setCurrentNum, setCurrent } = props;
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [drawing, setDrawing] = useState(false);
  const [hasContent, setHasContent] = useState(false);
  if (!question) return null;

  const getPos = (e: React.MouseEvent | React.TouchEvent) => {
    const canvas = canvasRef.current!;
    const rect = canvas.getBoundingClientRect();
    if ('touches' in e) {
      return { x: e.touches[0].clientX - rect.left, y: e.touches[0].clientY - rect.top };
    }
    return { x: e.clientX - rect.left, y: e.clientY - rect.top };
  };

  const startDraw = useCallback((e: React.MouseEvent | React.TouchEvent) => {
    e.preventDefault();
    const ctx = canvasRef.current?.getContext('2d');
    if (!ctx) return;
    setDrawing(true);
    setHasContent(true);
    ctx.beginPath();
    const pos = getPos(e);
    ctx.moveTo(pos.x, pos.y);
  }, []);

  const draw = useCallback((e: React.MouseEvent | React.TouchEvent) => {
    e.preventDefault();
    if (!drawing) return;
    const ctx = canvasRef.current?.getContext('2d');
    if (!ctx) return;
    const pos = getPos(e);
    ctx.lineTo(pos.x, pos.y);
    ctx.strokeStyle = '#000';
    ctx.lineWidth = 2;
    ctx.lineCap = 'round';
    ctx.stroke();
  }, [drawing]);

  const endDraw = useCallback(() => {
    setDrawing(false);
  }, []);

  const clearCanvas = useCallback(() => {
    const canvas = canvasRef.current;
    if (canvas) {
      const ctx = canvas.getContext('2d');
      ctx?.clearRect(0, 0, canvas.width, canvas.height);
      setHasContent(false);
    }
  }, []);

  const onDone = () => {
    if (!hasContent) {
      message.warning('请先签名');
      return;
    }
    const canvas = canvasRef.current;
    if (canvas) {
      const dataUrl = canvas.toDataURL('image/png');
      addRespondent({
        surveyId,
        type: question.type,
        questionId: question.id,
        value: [dataUrl],
        sn: generateRandom,
      });
    }
    handleJump(question, 'signed', setCurrent);
  };

  return (
    <>
      <h3>{question.serial ? question.serial + '-' : ''}{question.content}</h3>
      <Form.Item
        name={['question', "'" + question.id + "'"]}
        rules={[{ required: question.required === 1, message: '请签名' }]}
      >
        <div style={{ border: '1px solid #d9d9d9', borderRadius: 6, padding: 8, display: 'inline-block' }}>
          <canvas
            ref={canvasRef} width={320} height={160}
            style={{ background: '#fff', cursor: 'crosshair', display: 'block' }}
            onMouseDown={startDraw} onMouseMove={draw}
            onMouseUp={endDraw} onMouseLeave={endDraw}
            onTouchStart={startDraw} onTouchMove={draw} onTouchEnd={endDraw}
          />
        </div>
        <div style={{ marginTop: 8 }}>
          <Button icon={<ClearOutlined />} onClick={clearCanvas} size="small">清除</Button>
          <Button type="primary" onClick={onDone} size="small" style={{ marginLeft: 8 }}>确认签名</Button>
        </div>
      </Form.Item>
      <QJumpRules
        surveyId={surveyId} question={question} generateRandom={generateRandom}
        addRespondent={addRespondent} setCurrentNum={setCurrentNum} setCurrent={setCurrent}
        value={hasContent ? 'signed' : undefined}
      />
    </>
  );
};

export default QSignature;

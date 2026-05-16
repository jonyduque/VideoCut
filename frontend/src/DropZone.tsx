import React, { useState, useEffect } from 'react';
import { useStore } from './store';
import { SelectFile, GetVideoDuration } from '../wailsjs/go/main/App';

const DropZone: React.FC = () => {
  const [isDragging, setIsDragging] = useState(false);
  const { setFilePath, setDuration, setPart1Name, setPart2Name, setIsProcessing } = useStore();

  const handleFile = async (path: string) => {
    if (!path) return;
    
    setFilePath(path);
    const fileName = path.split(/[\\/]/).pop() || '';
    const baseName = fileName.substring(0, fileName.lastIndexOf('.')) || fileName;
    
    setPart1Name(`${baseName}_Parte1`);
    setPart2Name(`${baseName}_Parte2`);

    try {
      setIsProcessing(true);
      const duration = await GetVideoDuration(path);
      setDuration(duration);
    } catch (err) {
      console.error("Failed to get duration:", err);
      alert("Erro ao ler duração do vídeo. Verifique se o FFmpeg está instalado.");
    } finally {
      setIsProcessing(false);
    }
  };

  const onDrop = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragging(false);
    
    // Wails on Windows provides the path in different ways depending on version/config
    // Usually it's in dataTransfer.files[0].path
    const file = e.dataTransfer.files[0];
    if (file) {
      // @ts-ignore - path is present in Wails environment
      handleFile(file.path || '');
    }
  };

  const openDialog = async () => {
    try {
      const path = await SelectFile();
      if (path) {
        handleFile(path);
      }
    } catch (err) {
      console.error("SelectFile error:", err);
    }
  };

  return (
    <div 
      className={`drop-zone ${isDragging ? 'dragging' : ''}`}
      onDragOver={(e) => { e.preventDefault(); setIsDragging(true); }}
      onDragLeave={() => setIsDragging(false)}
      onDrop={onDrop}
      onClick={openDialog}
    >
      <div className="drop-content">
        <span className="icon">📁</span>
        <p>Arraste um vídeo aqui ou clique para selecionar</p>
      </div>
    </div>
  );
};

export default DropZone;

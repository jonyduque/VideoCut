import React, { useState, useEffect } from 'react';
import { useStore } from './store';
import { SelectFile, GetVideoDuration } from '../wailsjs/go/main/App';
import { OnFileDrop } from '../wailsjs/runtime/runtime';

const DropZone: React.FC = () => {
  const [isDragging, setIsDragging] = useState(false);
  const { setFilePath, setDuration, setPart1Name, setPart2Name, setIsProcessing, setCutPoint } = useStore();

  const handleFile = async (path: string) => {
    console.log("DropZone handleFile called with path:", path);
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
      setCutPoint(duration / 2); // Default to middle
    } catch (err) {
      console.error("Failed to get duration:", err);
      alert("Erro ao ler duração do vídeo. Verifique se o FFmpeg está instalado.");
    } finally {
      setIsProcessing(false);
    }
  };

  useEffect(() => {
    console.log("DropZone: Initializing OnFileDrop listener");
    
    // In Wails v2, OnFileDrop(callback, useDropTarget)
    // 'true' means it only fires for elements with --wails-drop-target: drop
    OnFileDrop((x, y, paths) => {
      console.log(`OnFileDrop fired at (${x},${y}) with paths:`, paths);
      if (paths && paths.length > 0) {
        handleFile(paths[0]);
      }
    }, true);

    // No direct unsubscribe for OnFileDrop in v2 runtime usually, 
    // but registering a new one replaces the old one.
  }, []);

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
      className="drop-zone"
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

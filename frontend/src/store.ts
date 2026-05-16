import { create } from 'zustand'

interface AppState {
  filePath: string
  duration: number
  cutPoint: number
  overlap: number
  outputDir: string
  part1Name: string
  part2Name: string
  useSourceFolder: boolean
  isProcessing: boolean
  ffmpegVersion: string
  
  setFilePath: (path: string) => void
  setDuration: (duration: number) => void
  setCutPoint: (point: number) => void
  setOverlap: (overlap: number) => void
  setOutputDir: (dir: string) => void
  setPart1Name: (name: string) => void
  setPart2Name: (name: string) => void
  setUseSourceFolder: (use: boolean) => void
  setIsProcessing: (is: boolean) => void
  setFfmpegVersion: (version: string) => void
}

export const useStore = create<AppState>((set) => ({
  filePath: '',
  duration: 0,
  cutPoint: 0,
  overlap: 5,
  outputDir: '',
  part1Name: '',
  part2Name: '',
  useSourceFolder: true,
  isProcessing: false,
  ffmpegVersion: '',

  setFilePath: (filePath) => set({ filePath }),
  setDuration: (duration) => set({ duration }),
  setCutPoint: (cutPoint) => set({ cutPoint }),
  setOverlap: (overlap) => set({ overlap }),
  setOutputDir: (outputDir) => set({ outputDir }),
  setPart1Name: (part1Name) => set({ part1Name }),
  setPart2Name: (part2Name) => set({ part2Name }),
  setUseSourceFolder: (useSourceFolder) => set({ useSourceFolder }),
  setIsProcessing: (isProcessing) => set({ isProcessing }),
  setFfmpegVersion: (ffmpegVersion) => set({ ffmpegVersion }),
}))

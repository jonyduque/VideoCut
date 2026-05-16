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
  successMessage: string
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
  setSuccessMessage: (msg: string) => void
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
  successMessage: '',
  ffmpegVersion: '',

  setFilePath: (filePath) => set({ filePath, successMessage: '' }),
  setDuration: (duration) => set({ duration }),
  setCutPoint: (cutPoint) => set({ cutPoint, successMessage: '' }),
  setOverlap: (overlap) => set({ overlap, successMessage: '' }),
  setOutputDir: (outputDir) => set({ outputDir, successMessage: '' }),
  setPart1Name: (part1Name) => set({ part1Name, successMessage: '' }),
  setPart2Name: (part2Name) => set({ part2Name, successMessage: '' }),
  setUseSourceFolder: (useSourceFolder) => set({ useSourceFolder, successMessage: '' }),
  setIsProcessing: (isProcessing) => set({ isProcessing }),
  setSuccessMessage: (successMessage) => set({ successMessage }),
  setFfmpegVersion: (ffmpegVersion) => set({ ffmpegVersion }),
}))

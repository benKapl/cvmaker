import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { CVData, PersonalInfo, Experience, Education, Skill } from '../lib/types';

interface CVStore {
  // État du CV en cours d'édition
  currentCV: CVData;
  
  // Actions pour les informations personnelles
  updatePersonalInfo: (info: Partial<PersonalInfo>) => void;
  
  // Actions pour les expériences
  addExperience: (experience: Omit<Experience, 'id'>) => void;
  updateExperience: (id: string, experience: Partial<Experience>) => void;
  removeExperience: (id: string) => void;
  
  // Actions pour l'éducation
  addEducation: (education: Omit<Education, 'id'>) => void;
  updateEducation: (id: string, education: Partial<Education>) => void;
  removeEducation: (id: string) => void;
  
  // Actions pour les compétences
  addSkill: (skill: Omit<Skill, 'id'>) => void;
  updateSkill: (id: string, skill: Partial<Skill>) => void;
  removeSkill: (id: string) => void;
  
  // Actions globales
  updateSummary: (summary: string) => void;
  resetCV: () => void;
  loadCV: (cvData: CVData) => void;
}

const initialCVData: CVData = {
  personalInfo: {
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    address: '',
    linkedIn: '',
    website: '',
  },
  experiences: [],
  education: [],
  skills: [],
  summary: '',
};

export const useCVStore = create<CVStore>()(
  persist(
    (set, get) => ({
      currentCV: initialCVData,
      
      updatePersonalInfo: (info) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            personalInfo: { ...state.currentCV.personalInfo, ...info },
          },
        })),
      
      addExperience: (experience) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            experiences: [
              ...state.currentCV.experiences,
              { ...experience, id: crypto.randomUUID() },
            ],
          },
        })),
      
      updateExperience: (id, experience) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            experiences: state.currentCV.experiences.map((exp) =>
              exp.id === id ? { ...exp, ...experience } : exp
            ),
          },
        })),
      
      removeExperience: (id) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            experiences: state.currentCV.experiences.filter((exp) => exp.id !== id),
          },
        })),
      
      addEducation: (education) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            education: [
              ...state.currentCV.education,
              { ...education, id: crypto.randomUUID() },
            ],
          },
        })),
      
      updateEducation: (id, education) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            education: state.currentCV.education.map((edu) =>
              edu.id === id ? { ...edu, ...education } : edu
            ),
          },
        })),
      
      removeEducation: (id) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            education: state.currentCV.education.filter((edu) => edu.id !== id),
          },
        })),
      
      addSkill: (skill) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            skills: [
              ...state.currentCV.skills,
              { ...skill, id: crypto.randomUUID() },
            ],
          },
        })),
      
      updateSkill: (id, skill) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            skills: state.currentCV.skills.map((s) =>
              s.id === id ? { ...s, ...skill } : s
            ),
          },
        })),
      
      removeSkill: (id) =>
        set((state) => ({
          currentCV: {
            ...state.currentCV,
            skills: state.currentCV.skills.filter((s) => s.id !== id),
          },
        })),
      
      updateSummary: (summary) =>
        set((state) => ({
          currentCV: { ...state.currentCV, summary },
        })),
      
      resetCV: () => set({ currentCV: initialCVData }),
      
      loadCV: (cvData) => set({ currentCV: cvData }),
    }),
    {
      name: 'cvmaker-storage', // Clé dans localStorage
      partialize: (state) => ({ currentCV: state.currentCV }), // Ne persister que le CV
    }
  )
);

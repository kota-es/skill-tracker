export type SearchedUser = {
  user_id: number;
  email: string;
  firstname: string;
  lastname: string;
  firstname_kana: string;
  lastname_kana: string;
  role: string;
  profile: Profile;
  skills: Skills[];
  lastUpdatedSkillDate?: string;
};

export type Profile = {
  id: number;
  user_id: number;
  notes: string;
  desires: string;
  dislikes: string;
  created_at: string;
  updated_at: string;
};

export type Skills = {
  id: number;
  skill_id: number;
  user_id: number;
  level: number;
  interested: boolean;
  created_at: string;
  updated_at: string;
};

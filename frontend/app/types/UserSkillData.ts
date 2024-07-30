export type UserSkillData = SkillDataInCategory[];

type SkillDataInCategory = {
  id: number;
  name: string;
  skills: UserSKills[];
};

export type UserSKills = {
  id: number;
  name: string;
  description: string;
  level: number;
  levelExplanation: string;
  interested: boolean;
};

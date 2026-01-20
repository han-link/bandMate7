export interface APIResponse<Type> {
    data: Type;
    error: string;
    success: boolean;
}

interface Base {
    id: string;
    createdAt: string;
    updatedAt: string;
}

export interface Performance extends Base {
    bpm: number;
    name: string;
    cover: Resource;
}

export interface Resource extends Base {
    filename: string;
    role: UserRole;
    type: ResourceType;
}

interface UserRole extends Base {
    name: string;
}

// @ts-ignore
export enum ResourceType {
    IMAGE = "image",
    VIDEO = "video",
    AUDIO = "audio",
    DOCUMENT = "document"
}
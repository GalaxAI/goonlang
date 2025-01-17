import axiosInstance from '../config/axiosConfig';
import { AxiosResponse } from 'axios';
import { Base64ImagesResponse, RuneMatrix3D } from '../types/images';
interface ImageResponse {
  imageResponse: Base64ImagesResponse;
  asciiArt: RuneMatrix3D;
}

interface EdgeDetectionParams {
  base64Image: string;
  gradientThreshold: number;
  threshold: number;
  tau: number;
}

export const loadImage = (params: EdgeDetectionParams): Promise<ImageResponse> => {
  console.log('loadImage function called');
  return axiosInstance.post<ImageResponse>('/edge-detect-ascii', params)
    .then((response: AxiosResponse<ImageResponse>) => {
      console.log('Response data:', response.data);
      return response.data;
    })
    .catch((error: Error) => {
      console.error('Error loading image:', error);
      throw error;
    });
};

export const colorDownsample = (base64Image: string): Promise<ImageResponse> => {
  console.log('loadImageColor function called');
  return axiosInstance.post<ImageResponse>('/color-downsample', {
    base64Image
  })
    .then((response: AxiosResponse<ImageResponse>) => {
      return response.data;
    })
    .catch((error: Error) => {
      console.error('Error loading color image:', error);
      throw error;
    });
};
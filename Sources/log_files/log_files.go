package log_files

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client LogFilesServiceClient
}
    /*
         Get List of log files.

         

         Returns
         -------
         True
         Entries : []*Entry
              List of entries

         
    */


    func(s *ServiceImpl)GetEntries(ctx context.Context, ) (*GetEntriesResponse, error){
        request := &GetEntriesRequest{}
    	response, err := s.Client.GetEntries(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Download log file.

         Parameters
         ----------
         entry *Entry , path string
    */

    func (a *ServiceImpl) DownloadLogFile(ctx context.Context, entry *Entry , path string) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeDownloadLogFileRequest{}
    		request.Entry = entry
            	
        	request.Path = path
        	stream, err := a.Client.SubscribeDownloadLogFile(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &DownloadLogFileResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					break
    				}
    				if err != nil {
    					fmt.Printf("Unable to receive message %v", err)
    					break
    				}
    				ch <- m.GetProgress()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Download log file synchronously.

         Parameters
         ----------
         entry *Entry 
            path string

         Returns
         -------
         False
         Progress : ProgressData
              Progress if result is progress

         
    */


    func(s *ServiceImpl)DownloadLogFile(ctx context.Context, entry *Entry , path string) (*DownloadLogFileResponse, error){
        request := &DownloadLogFileRequest{}
    	request.Entry = entry
            
        request.Path = path
        response, err := s.Client.DownloadLogFile(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Erase all log files.

         
    */

    func(s *ServiceImpl)EraseAllLogFiles(ctx context.Context, )(*EraseAllLogFilesResponse, error){
        
        request := &EraseAllLogFilesRequest{}
    	response, err := s.Client.EraseAllLogFiles(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
package main

import (
	"INNLookupPro/cmd/logger"
	"INNLookupPro/protos"
	"context"
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	protos.UnimplementedCompanyInfoServiceServer
}

const (
	defaultState = "Not found"
)

func (s *Server) GetCompanyInfo(ctx context.Context, req *protos.CompanyInfoRequest) (*protos.CompanyInfoResponse, error) {
	inn := req.Inn

	kpp := defaultState
	companyName := defaultState
	directorName := defaultState

	c := colly.NewCollector(colly.AllowedDomains("www.rusprofile.ru"))

	c.OnHTML(".company-item", func(e *colly.HTMLElement) {
		warningText := e.ChildText("div.company-item-status span.warning-text")
		tempInnElement := e.DOM.Find(".company-item-info > dl > dt:contains('ИНН') + dd")
		tempInn := tempInnElement.Text()
		if warningText == "" && inn == tempInn {
			// Only visit the link if there's no warning and inn is the real one
			link := e.ChildAttr(".company-item__title a", "href")
			// Visit link found on page
			err := c.Visit(e.Request.AbsoluteURL(link))
			if err != nil {
				logger.Println("Server error while scraping:", err)
				kpp = defaultState
				companyName = defaultState
				directorName = defaultState
				return
			}
		}
	})

	c.OnHTML("h2.company-name", func(e *colly.HTMLElement) {
		companyName = e.Text
	})

	c.OnHTML("#clip_kpp", func(e *colly.HTMLElement) {
		kpp = e.Text
	})

	c.OnHTML("div.company-row.hidden-parent span.company-info__text a span", func(e *colly.HTMLElement) {
		directorName = e.Text
	})

	err := c.Visit("https://www.rusprofile.ru/search?query=" + inn)
	if err != nil {
		logger.Println("Server error while scraping:", err)
		return nil, err
	}

	response := &protos.CompanyInfoResponse{
		Inn:          inn,
		Kpp:          kpp,
		CompanyName:  companyName,
		DirectorName: directorName,
	}

	return response, nil
}

func main() {
	portPtr := flag.String("Port", "8080", "Port to listen on")
	serverHelpPtr := flag.Bool("help", false, "Show help message for the server")
	flag.Parse()

	if *serverHelpPtr {
		flag.Usage()
		return
	}

	listenAddr := ":" + *portPtr

	s := grpc.NewServer()
	srv := &Server{}
	protos.RegisterCompanyInfoServiceServer(s, srv)

	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("gRPC server is running on", listenAddr)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

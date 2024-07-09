package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "pdfGen/documentgeneration"
    "pdfGen/controllers/payloads"
)

func CreateMemarts(c *gin.Context) {
    var body payloads.CreateMemartRequest
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    data := documentgeneration.PageData{
        Company: documentgeneration.Company{
            Name:         body.Company.Name,
            Office:       body.Company.Office,
            Objectives:      body.Company.Objectives,
            Liability:    body.Company.Liability,
            ShareCapital: body.Company.ShareCapital,
        },
        Subscribers: []documentgeneration.Subscriber{},
        Date:       documentgeneration.Date{
            Day:           body.Date.Day,
            Month:         body.Date.Month,
            Year:          body.Date.Year,
        },
        CheckOption:  documentgeneration.CheckOption{
            AdoptTableII: body.CheckOption.AdoptTableII,
            AdoptTableIIWithModification: body.CheckOption.AdoptTableIIWithModification,
        },
    }

    // Map subscribers
    for _, subscriber := range body.Subscribers {
        data.Subscribers = append(data.Subscribers, documentgeneration.Subscriber{
            Name:       subscriber.Name,
            Occupation: subscriber.Occupation,
            Shares:     subscriber.Shares,
            Signature:  subscriber.Signature,
        })
    }

    err := documentgeneration.GeneratePDF(data, "./output/MEMARTS.pdf")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "PDF generated successfully"})
}

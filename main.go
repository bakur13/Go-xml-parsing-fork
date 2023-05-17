
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type SalesOrderItem struct {
	CreatWho       string `xml:"CREAT_WHO"`
	ID             int    `xml:"SALES_ORDER_ITEM_ID"`
	ActionDate     string `xml:"ACTION_DATE"`
	ReasonCode     string `xml:"REASON_CODE"`
	CustomerID     int    `xml:"CUSTOMER_ACCOUNT_KEY>CUSTOMER_ACCOUNT_ID"`
	ServiceID      int    `xml:"SERVICE>SERVICE_ID"`
	ServiceName    string `xml:"SERVICE>SERVICE_NAME"`
	Description    string `xml:"SERVICE>DESCRIPTION"`
	ServiceSpecID  int    `xml:"SERVICE>SERVICE_SPECIFICATION_ID"`
	NetworkService string `xml:"SERVICE>NETWORK_SERVICE_CODE"`
	ProductID      int    `xml:"PRODUCT>PRODUCT_KEY>PRODUCT_ID"`
	ProductName    string `xml:"PRODUCT>PRODUCT_DEFINITION>PRODUCT_NAME"`
	ProductDesc    string `xml:"PRODUCT>PRODUCT_DEFINITION>DESCRIPTION"`
	ProductStatus  string `xml:"PRODUCT>PRODUCT_DEFINITION>PRODUCT_STATUS"`
	ProductNumber  int    `xml:"PRODUCT>PRODUCT_DEFINITION>PRODUCT_NUMBER"`
	AccountTaxInfo int    `xml:"PRODUCT>PRODUCT_DEFINITION>ACCOUNT_TAX_INFO"`
	UOMTypeID      int    `xml:"PRODUCT>PRODUCT_PRICE_INFO>UOM_TYPE_ID"`
	TariffID       int    `xml:"PRODUCT>PRODUCT_PRICE_INFO>TARIFF_ID"`
	MacID          string `xml:"PRODUCT>PRODUCT_CHAR_VALUE>PRODUCT_CHAR_VALUE"`
	BillingInd     string `xml:"PRODUCT>PRODUCT_CHAR_VALUE>BILLING_INDICATOR"`
	ResourceInd    string `xml:"PRODUCT>PRODUCT_CHAR_VALUE>RESOURCE_INDICATOR"`
	OfferID        int    `xml:"OFFER>PRODUCT_OFFER_KEY>PRODUCT_OFFER_ID"`
	OfferName      string `xml:"OFFER>PRODUCT_OFFER_DEFINITION>OFFER_NAME"`
	OfferStatus    string `xml:"OFFER>PRODUCT_OFFER_DEFINITION>OFFER_STATUS"`
	IsBundle       string `xml:"OFFER>PRODUCT_OFFER_DEFINITION>IS_BUNDLE"`
	BillingAccID   int    `xml:"BILLING_INFO>BILLING_ACCOUNT_KEY>BILLING_ACCOUNT_ID"`
	Quantity       int    `xml:"QUANTITY"`
}

type Product struct {
	XMLName         xml.Name         `xml:"product"`
	SalesOrderItems []SalesOrderItem `xml:"SALES_ORDER_ITEM"`
}

func main() {
	xmlFile, err := os.Open("product.xml")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, _ := ioutil.ReadAll(xmlFile)

	var p Product
	xml.Unmarshal(xmlData, &p)
	fmt.Println("Sales Order Items:")
	for _, item := range p.SalesOrderItems {
		fmt.Println("----------------------------------")
		fmt.Printf("CreatWho: %s\n", item.CreatWho)
		fmt.Printf("ID: %d\n", item.ID)
		fmt.Printf("Action Date: %s\n", item.ActionDate)
		fmt.Printf("Reason Code: %s\n", item.ReasonCode)
		fmt.Printf("Customer Account ID: %d\n", item.CustomerID)
		fmt.Printf("Service ID: %d\n", item.ServiceID)
		fmt.Printf("Service Name: %s\n", item.ServiceName)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("Service Specification ID: %d\n", item.ServiceSpecID)
		fmt.Printf("Network Service Code: %s\n", item.NetworkService)
		fmt.Printf("Product ID: %d\n", item.ProductID)
		fmt.Printf("Product Name: %s\n", item.ProductName)
		fmt.Printf("Product Description: %s\n", item.Description)
		fmt.Printf("Product Status: %s\n", item.ProductStatus)
		fmt.Printf("Product Number: %d\n", item.ProductNumber)
		fmt.Printf("Account Tax Info: %d\n", item.AccountTaxInfo)
		fmt.Printf("UOM Type ID: %d\n", item.UOMTypeID)
		fmt.Printf("Tariff ID: %d\n", item.TariffID)
		fmt.Printf("Mac ID: %s\n", item.MacID)
		fmt.Printf("Billing Indicator: %s\n", item.BillingInd)
		fmt.Printf("Resource Indicator: %s\n", item.ResourceInd)
		fmt.Printf("Offer ID: %d\n", item.OfferID)
		fmt.Printf("Offer Name: %s\n", item.OfferName)
		fmt.Printf("Offer Status: %s\n", item.OfferStatus)
		fmt.Printf("Is Bundle: %s\n", item.IsBundle)
		fmt.Printf("Billing Account ID: %d\n", item.BillingAccID)
		fmt.Printf("Quantity: %d\n", item.Quantity)
		fmt.Println("----------------------------------")
	}

}

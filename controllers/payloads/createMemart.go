package payloads

type CreateMemartRequest struct {
    Company     CompanyPayload     `json:"company"`
    Subscribers []SubscriberPayload `json:"subscribers"`
    Date       DatePayload       `json:"date"`
    CheckOption CheckOptionPayload
}

type CheckOptionPayload struct{
    AdoptTableII                 bool `json:"adopt_table_a_part_II"`
    AdoptTableIIWithModification bool `json:"adopt_table_a_part_II_with_modification"`
}

type CompanyPayload struct {
    Name         string `json:"name"`
    Office       string `json:"office_location"`
    Objectives   []string `json:"objectives"`
    Liability    string `json:"liability"`
    ShareCapital string `json:"share_amount"`
}

type SubscriberPayload struct {
    Name       string `json:"name"`
    Occupation string `json:"occupation"`
    Shares     string `json:"shares"`
    Signature  string `json:"signature"`
}

// type ObjectivesPayload struct {
//     Objective string
// }

type DatePayload struct {
    Day string      `json:"day"`
    Month string     `json:"month"`
    Year string     `json:"year"`
}

# Email Verifier

A command-line tool written in Go that verifies domain email configurations by checking MX records, SPF, and DMARC settings.

## Features

- Checks for MX (Mail Exchange) records
- Verifies SPF (Sender Policy Framework) records
- Validates DMARC (Domain-based Message Authentication, Reporting, and Conformance) records
- Outputs results in CSV format

## Installation

```bash
git clone github.com/fbdaf/email-verifier
cd email-verifier
go build
```

## Usage

The tool reads domain names from standard input and outputs the verification results in CSV format. Each line of output contains:
- Domain name
- MX record status
- SPF record status
- SPF record content
- DMARC record status
- DMARC record content

### Example

```bash
echo "google.com" | ./email-verifier
```

Output format:
```csv
domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord
google.com, true, true, v=spf1 include:_spf.google.com ~all, true, v=DMARC1 p=reject
```

### Batch Processing

You can verify multiple domains by providing them in a file:

```bash
cat domains.txt | ./email-verifier > results.csv
```

## Error Handling

The tool handles DNS lookup errors gracefully and logs them while continuing to process remaining domains. Errors are logged to stderr while the CSV output goes to stdout.

## Requirements

- Go 1.21.4 or higher
- Internet connection for DNS lookups

## Acknowledgments

- Go DNS package documentation
- Email authentication standards (SPF, DMARC)
